package scan_tester

import "fmt"

func ErrorResponseSpec(code string) string {
	return fmt.Sprintf(
		`info:
  title: Test
paths:
  /api/document:
    post:
      operationId: postDocument
      responses:
        "200":
          description: Stored.
        "%s":
          description: Failed.
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ErrorResponse"
components:
  schemas:
    ErrorResponse:
      type: object
`,
		code,
	)
}
