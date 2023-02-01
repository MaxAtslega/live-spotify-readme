# Live-Spotify-Template

## Review
[<img src="https://spotify.atslega.dev/api/spotify" alt="Spotify Playing" width="380"/>](https://github.com/MaxAtslega/live-spotify-readme)

## How to get the SPOTIFY_REFRESH_TOKEN?
1. Create a [Spotify Application](https://developer.spotify.com/dashboard/applications)
2. Take a note of the Client ID, and Client Secret for setting up .env
3. Edit settings and add `http://localhost/callback/` as your Redirect URI;
4. For request authorization (log in and authorize access) navigate to: `https://accounts.spotify.com/authorize?client_id={SPOTIFY_CLIENT_ID}&response_type=code&scope=user-read-currently-playing&redirect_uri=http://localhost/callback/`. You will be redirected back to your specified redirect uri. The response query string should contains an authorization CODE (e.g. https://example.com/callback?code={CODE}); 
5. Create a string combining `{SPOTIFY_CLIENT_ID}:{SPOTIFY_CLIENT_SECRET}` (e.g. 6hg43g34g43g34g43o45u2g34f3j2:kj2g3v2j4j2j4hj2v4hj2v34) and encode into Base64.
6. Run the following curl command to: `curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -H "Authorization: Basic {BASE64}" -d "grant_type=authorization_code&redirect_uri=http://localhost/callback/&code={CODE}" https://accounts.spotify.com/api/token`. This will return a JSON response containing the SPOTIFY_REFRESH_TOKEN.
7. Save the Refresh token.
