package handlers

// import (
// 	"errors"
// 	"github/jaseelaali/orchid/repository"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// )

// // Define a test error to use in the tests
// var someError = errors.New("some error")

// // Define a struct to hold the test cases
// type testCase struct {
// 	name       string // Name of the test case
// 	queryParam string // Query parameter to pass to the function
// 	mockError  error  // Error to return from the mock repository function
// 	statusCode int    // Expected HTTP status code of the response
// 	response   string // Expected response body of the response
// }

// // Define the test function
// func TestReturnStatus(t *testing.T) {
// 	// Set Gin to test mode
// 	gin.SetMode(gin.TestMode)

// 	// Define the test cases
// 	testCases := []testCase{
// 		{
// 			name:       "Success",
// 			queryParam: "payment123",
// 			mockError:  nil,
// 			statusCode: http.StatusOK,
// 			response:   `{"message":"delivery completed"}`,
// 		},
// 		{
// 			name:       "Missing Payment ID",
// 			queryParam: "",
// 			mockError:  nil,
// 			statusCode: http.StatusBadRequest,
// 			response:   `{"message":"didn't get payment id"}`,
// 		},
// 		{
// 			name:       "Error Updating Data",
// 			queryParam: "payment456",
// 			mockError:  someError,
// 			statusCode: http.StatusBadRequest,
// 			response:   `{"message":"couldn't update the data"}`,
// 		},
// 	}

// 	// Loop over the test cases
// 	for _, tc := range testCases {
// 		// Create a new Gin context for each test case
// 		context, _ := gin.CreateTestContext(httptest.NewRecorder())

// 		// Set the query parameter for this test case
// 		context.Request = httptest.NewRequest(http.MethodGet, "/endpoint?paymentid="+tc.queryParam, nil)

// 		// Mock the repository function
// 	// 	if tc.mockError != nil {
// 	// 		repository.ReturnStatusChange = func(paymentID string) error {
// 	// 			return tc.mockError
// 	// 		}
// 	// 	} else {
// 	// 		repository.ReturnStatusChange = func(paymentID string) error {
// 	// 			return nil
// 	// 		}
// 	// 	}

// 	// 	// Call the function and check the response
// 	// 	ReturnStatus(context)

// 	// 	if context.Writer.Status() != tc.statusCode {
// 	// 		t.Errorf("Expected status code %d but got %d", tc.statusCode, context.Writer.Status())
// 	// 	}

// 	// 	if body := context.Writer.Body.String(); body != tc.response {
// 	// 		t.Errorf("Expected response body %s but got %s", tc.response, body)
// 	// 	}
// 	// }
// }
