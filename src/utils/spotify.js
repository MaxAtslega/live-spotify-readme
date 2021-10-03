import fetch from "node-fetch";
import {stringify} from "querystring";

async function getAuthorizationToken() {

  const client_id = process.env.SPOTIFY_CLIENT_ID;
  const client_secret = process.env.SPOTIFY_CLIENT_SECRET;
  const refresh_token = process.env.SPOTIFY_REFRESH_TOKEN;

  const basic = Buffer.from(`${client_id}:${client_secret}`).toString("base64");

  const url = new URL("https://accounts.spotify.com/api/token");

  const body = stringify({
    grant_type: "refresh_token",
    refresh_token,
  });

  const response = await fetch(`${url}`, {
    method: "POST",
    headers: {
      Authorization : `Basic ${basic}`,
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body,
  }).then((r) => r.json());



  return `Bearer ${response["access_token"]}`;
}

export async function nowPlaying() {
  const Authorization = await getAuthorizationToken();

  const response = await fetch("https://api.spotify.com/v1/me/player/currently-playing", {
    headers: {
      Authorization,
    },
  });

  const status = response["status"];
  if (status === 204) {
    return {};
  } else if (status === 200) {

    const data = await response.json();
    return data;
  }
}




