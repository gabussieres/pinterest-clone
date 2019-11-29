/* global process */
/* global module */

const URL = "https://morning-dawn-44308.herokuapp.com/";

function checkStatus(response) {
  if (response.ok) {
    return response;
  } else if (response.status === 401) {
    // Handle sso redirect
    window.location = response.headers.get("auth-redirect");
    return undefined;
  }
  const error = new Error(`HTTP Error ${response.statusText}`);
  error.status = response.statusText;
  error.response = response;
  console.error(error); // eslint-disable-line no-console
  throw error;
}

function parseJSON(response) {
  return response.json();
}

async function getResource(path) {
  return fetch(path, { method: "GET" })
    .then(checkStatus)
    .then(parseJSON);
}

// /delete_pin should be a delete operation on /pin, but needed to work around CORS,
// so for our purposes I'm using a  a temporary GET endpoint to handle the deletion
// this is a function for the exclusive use of delete_pin
async function getResourceForDeletePin(path) {
  return fetch(path, { method: "GET" }).then(checkStatus);
}

async function fetchUser(user_id) {
  const response = await getResource(`${URL}/user?user_id=${user_id}`);
  return response;
}

async function fetchPinDetails(pin_id) {
  const response = await getResource(`${URL}/pin?pin_id=${pin_id}`);
  return response;
}

async function fetchUserPins(user_id, pin_id) {
  const response = await getResource(
    `${URL}/pins?user_id=${user_id}&pin_id=${pin_id}`
  );
  return response;
}

async function fetchFeed(pinAmount) {
  const response = await getResource(`${URL}/feed?pin_amount=${pinAmount}`);
  return response;
}

async function deletePin(user_id, pin_id) {
  const response = await getResourceForDeletePin(
    `${URL}/delete_pin?user_id=${user_id}&pin_id=${pin_id}`
  );
  return response;
}

const API = { fetchFeed, fetchPinDetails, fetchUser, fetchUserPins, deletePin };

export default API;
