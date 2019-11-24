package config

// Config holds all config variables
type Config struct {
	DynamoConfig DynamoConfig
}

// DynamoConfig holds config variable related to DynamoDB
type DynamoConfig struct {
	UserTable     TableConfig
	PinTable      TableConfig
	UserPinsTable TableConfig
}

// TableConfig holds the configuration for a single DynamoDB table
type TableConfig struct {
	Name         string
	PartitionKey string
	SortKey      string
}

/*
There are 3 tables in DynamoDB:

-- pinterest_users
PK: id (string)
Other Attributes: name, image_url, followers, following

-- pinterest_pins
PK: id (string)
Other Attributes: image_url, source_url, title, description, posted_by

-- pinterest_user_pins
This is a join table, that represents the 1:many relationship from user to pins
PK: user_id (string)
SK: pin_id (string)
*/

// PinterestConfig holds env config variables
var PinterestConfig = Config{
	DynamoConfig: DynamoConfig{
		UserTable: TableConfig{
			Name:         "pinterest_users",
			PartitionKey: "id",
		},
		PinTable: TableConfig{
			Name:         "pinterest_pins",
			PartitionKey: "id",
		},
		UserPinsTable: TableConfig{
			Name:         "pinterest_user_pins",
			PartitionKey: "user_id",
			SortKey:      "pin_id",
		},
	},
}
