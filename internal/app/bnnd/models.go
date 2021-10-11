package bnnd

type (
	// Actual model
	BananaDuck struct {
		ID    string
		Owner string
		// swagger:ignore
		Data interface{}
	}

	// Model for swagger
	// swagger:model bananaduckSwagger
	BananaDuckSwagger struct {
		// swagger:allOf
		BananaDuck
		// The specific data
		Data TheWayYouWouldWriteItInKBMF
	}

	// This model may or may not already exist
	TheWayYouWouldWriteItInKBMF struct {
		Val1 string
		Val2 string
		Val3 string
	}
)
