# pinterest-clone

## API

**/user**

Takes a UserID, returns the user model for that ID

**/pin**

Takes a PinID, returns the pin model for that ID

**/pins**

Takes a UserID, returns the pin models for that user

**/feed**

Returns a random array of pins

## Tables

There are 3 tables in DynamoDB:

**pinterest\_users**

PK: id (string)
Other Attributes: name, image\_url, followers, following

**pinterest\_pins**

PK: id (number)
Other Attributes: image\_url, source\_url, title, description, posted\_by

**pinterest\_user\_pins**

This is a join table, that represents the 1:many relationship from user to pins
PK: user\_id (string)
SK: pin\_id (number)

## Important considerations

There are a few things I would have added, but couldn't due to time constraints. This includes:

1. Deploying the API to a server

I would distribute the service over 2 ECS clusters (in case one goes down). However, I am interested in learning how to leverage EKS (Kubernetes) to manage microservices.

2. Building a deploy pipeline

I would use CircleCI for my deployment pipeline with the following flow

**Build -> Test -> Hold -> Deploy**

where,

1. Build: Codegen via go-swagger, then build the executable
2. Test: Run all unit tests
3. Hold: An extra step that must be manually unheld to proceed to deployment (useful for production)
4. Deploy: Deploy to server
