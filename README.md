# Svoya igra API

Svoya igra is an intellectual game.There will be specific topic and 5 questions related to the topic with 1-5 level of difficulty.

# Project description

In this project I created CRUD API which user can create new topic with questions,able to see the topics,change the questions on the topic and delete the topic.
Each question has anwer next to it.

# Usage of endpoints.

## Creating topics

User need to send `POST` request to `/topic` endpoint with JSON data including `title`,`question1`,`question2`,`question3`,`question4`,`question5`.
ID will be automatically added to the JSON data.

<img width="830" alt="Screenshot 2024-06-24 at 09 38 30" src="https://github.com/Muhammadjon1304/svoya-igra-api/assets/142288222/50ae54a5-71a1-4f78-9535-afe2ed9b26b7">
Request body must be like above.

## Getting all topics

User need to send `GET` request to `/topic` endpoint.Server returns the JSON data imcluding all topics.

## Getting one topic

User need to send `GET` request to `/topic/id` endpoint.

<img width="717" alt="Screenshot 2024-06-24 at 09 54 06" src="https://github.com/Muhammadjon1304/svoya-igra-api/assets/142288222/e1471130-0946-4377-a7b5-cae4d5d0bd10">

Server returns the JSON data with topic with requested ID.

## Updating the topics

User need to send `PUT` request to `/topic/id` endpoint with JSON data including `title`,`question1`,`question2`,`question3`,`question4`,`question5`.

<img width="720" alt="Screenshot 2024-06-24 at 10 11 06" src="https://github.com/Muhammadjon1304/svoya-igra-api/assets/142288222/1fc4754f-823a-46bc-975f-0e0d67f8fca6">

Server returns the JSON data with topic with updated content.

## Deleting the topics

User need to send `DELETE` request to `/topic/id` endpoint.

<img width="718" alt="Screenshot 2024-06-24 at 10 13 45" src="https://github.com/Muhammadjon1304/svoya-igra-api/assets/142288222/b20894b8-2cdd-4107-a63e-3068da621554">

Server deletes the topic which has requested ID.
