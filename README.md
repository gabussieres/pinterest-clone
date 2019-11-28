# pinterest-clone

**Running:**
To start the frontend:

1. cd client
2. npm i
3. npm start

To start the backend:
NOTE: the API is not hosted yet, so no data will appear on the frontend unless local dynamo is set up. I am currently working on hosting the backend on heroku.

1. Go to root directory
2. go get ./...
3. go install ./cmd/pinterest-clone-server
4. go run ./cmd/pinterest-clone-server/ --port 3001

# Frontend

The architecture comprises of React, Redux, reselect, sagas, and react-router.

## Important considerations

Here is a list of features that were not implemented or contain bugs due to contraints in the current implementation

- Due to time constraints, Github login was added, but not authentication. Instead, the client is assuming a logged in status for the user Gabriel.
- Local storage was not implemented, which means that state does not persist on location change (e.g. '/' to '/gabriel'). So, the user needs to log in on every location. This also leads to a bug where the "Delete Pin" button is visible on all pins, even though it only works on pins saved by the user. I would consider implementing a local storage solution, which could be done with the react-persist package.
- Again, due to time constraints I was not able to reach the URL shortening portion of the challenge.
- Finally, I had enough time to complete unit tests for the backend, but not for the frontend. I would have implemented them with Jest, as I have had experience builing 200-300 Jest unit tests for an internal dashboard at my previous company.

# Backend

## API

**/user**

Takes a UserID, returns the user model for that ID

**/pin**

Takes a PinID, returns the pin model for that ID

**/pins**

Takes a UserID, returns the pin models for that user

**/feed**

Takes a PinAmount (number of pins to return). Returns a random array of pins

**/delete_pin**

Takes a PinID, deletes the pin model for that ID

- Note: Since the API is not hosted, the frontend implementation for deleting pins required a workaround, because delete operations cannot be performed without CORS. Therefore, the DELETE method for pins was temporarily converted into a GET. Otherwise, it would be a DELETE on /pin/:pin_id

## Tables

There are 3 tables in DynamoDB:

**pinterest_users**

PK: id (string)
Other Attributes: name, image_url, followers, following

**pinterest_pins**

PK: id (number)
Other Attributes: image_url, source_url, title, description, posted_by

**pinterest_user_pins**

This is a join table, that represents the 1:many relationship from user to pins
PK: user_id (string)
SK: pin_id (string)

## Important considerations

There are a few things I would have added, but couldn't due to time constraints. This includes:

1. Deploying the API to a server

Currently working on deploying to heroku. Realistically, I would distribute the service over 2 ECS clusters (in case one goes down). However, I am interested in learning how to leverage EKS (Kubernetes) to manage microservices.

- Note: Since the API is not hosted, the frontend implementation for deleting pins required a workaround, since delete operations cannot be performed without CORS. Therefore, the DELETE method for pins was temporarily converted into a GET.

2. Building a deploy pipeline

I would use CircleCI for my deployment pipeline with the following flow

**Build -> Test -> Hold -> Deploy**

where,

1. Build: Codegen via go-swagger, then build the executable
2. Test: Run all unit tests
3. Hold: An extra step that must be manually unheld to proceed to deployment (useful for production)
4. Deploy: Deploy to server
