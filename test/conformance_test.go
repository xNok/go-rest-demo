package test

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/assert"
	"github.com/xNok/go-rest-demo/pkg/client"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	serverUrl = "http://localhost:8080"
)

var implementations = []struct {
	name string
	path string
}{
	{"standard-library", "../cmd/standardlib/main.go"},
	{"gorilla", "../cmd/gorilla/main.go"},
	{"gin", "../cmd/gin/main.go"},
}

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestConformance(t *testing.T) {
	for _, impl := range implementations {
		t.Run(impl.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", impl.path)
			err := cmd.Start()
			require.NoError(t, err, "Failed to start server")

			// It's not ideal to use a fixed sleep, but for this demo, it's sufficient
			// A better approach would be to poll the server's health check endpoint
			time.Sleep(2 * time.Second)

			defer func() {
				err := cmd.Process.Kill()
				require.NoError(t, err, "Failed to kill server process")
			}()

			runTests(t)
		})
	}
}

func runTests(t *testing.T) {
	c, err := client.NewClientWithResponses(serverUrl)
	require.NoError(t, err, "Failed to create client")

	// 1. Initially, there should be no recipes
	t.Run("list recipes initially", func(t *testing.T) {
		resp, err := c.GetRecipesWithResponse(context.Background())
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode())
		assert.Empty(t, resp.JSON200)
	})

	// 2. Create a new recipe
	recipeName := "Ham and Cheese Sandwich"
	ingredients := []client.Ingredient{
		{Name: string_ptr("Ham")},
		{Name: string_ptr("Cheese")},
		{Name: string_ptr("Bread")},
	}
	recipe := client.Recipe{
		Name:        string_ptr(recipeName),
		Ingredients: &ingredients,
	}
	t.Run("create a new recipe", func(t *testing.T) {
		resp, err := c.PostRecipesWithResponse(context.Background(), recipe)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode())
	})

	// 3. List recipes and expect one recipe
	recipeID := slug.Make(recipeName)
	t.Run("list recipes with one recipe", func(t *testing.T) {
		resp, err := c.GetRecipesWithResponse(context.Background())
		if err != nil {
			t.Logf("Error from GetRecipesWithResponse: %v", err)
		}
		if resp != nil {
			t.Logf("GetRecipesWithResponse status: %s", resp.Status())
			t.Logf("GetRecipesWithResponse body: %s", string(resp.Body))
			t.Logf("GetRecipesWithResponse headers: %v", resp.HTTPResponse.Header)
		}
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode())
		require.NotNil(t, resp.JSON200, "resp.JSON200 should not be nil")
		assert.Contains(t, *resp.JSON200, recipeID)
	})

	// 4. Get the created recipe
	t.Run("get created recipe", func(t *testing.T) {
		resp, err := c.GetRecipesIdWithResponse(context.Background(), recipeID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode())
		require.NotNil(t, resp.JSON200)
		assert.Equal(t, *recipe.Name, *resp.JSON200.Name)
	})

	// 5. Update the recipe
	updatedIngredients := []client.Ingredient{
		{Name: string_ptr("Ham")},
		{Name: string_ptr("Cheese")},
		{Name: string_ptr("Bread")},
		{Name: string_ptr("Butter")},
	}
	updatedRecipe := client.Recipe{
		Name:        string_ptr(recipeName),
		Ingredients: &updatedIngredients,
	}
	t.Run("update recipe", func(t *testing.T) {
		resp, err := c.PutRecipesIdWithResponse(context.Background(), recipeID, updatedRecipe)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode())
	})

	// 6. Get the updated recipe and check if it's updated
	t.Run("get updated recipe", func(t *testing.T) {
		resp, err := c.GetRecipesIdWithResponse(context.Background(), recipeID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode())
		require.NotNil(t, resp.JSON200)
		assert.Equal(t, *updatedRecipe.Name, *resp.JSON200.Name)
		assert.Len(t, *resp.JSON200.Ingredients, 4)
	})

	// 7. Delete the recipe
	t.Run("delete recipe", func(t *testing.T) {
		resp, err := c.DeleteRecipesIdWithResponse(context.Background(), recipeID)
		require.NoError(t, err)
		// The standard library implementation returns 200, but 204 is more appropriate
		// For the sake of conformance, we accept both
		assert.True(t, resp.StatusCode() == http.StatusOK || resp.StatusCode() == http.StatusNoContent)
	})

	// 8. Get the deleted recipe and expect a 404
	t.Run("get deleted recipe", func(t *testing.T) {
		resp, err := c.GetRecipesIdWithResponse(context.Background(), recipeID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode())
	})
}

func string_ptr(s string) *string {
	return &s
}
