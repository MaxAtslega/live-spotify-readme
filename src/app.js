import express from "express";
import bodyParser from "body-parser";
import logger from "morgan";

import Canvas, { registerFont }  from 'canvas'
require("dotenv").config();

import {fillMixedTextLeft, roundRect} from './utils'

import {nowPlaying} from "./utils/spotify"

let app = express();

//Set Port
const PORT = process.env.PORT || 3000;
app.set("port", PORT);


app.use(logger("tiny"));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

//Register Font
registerFont("./src/assets/fonts/Roboto/Roboto-Bold.ttf", { family: "RobotoBold" });
registerFont("./src/assets/fonts/Roboto/Roboto-Regular.ttf", { family: "RobotoRegular" });

app.get("/api/spotify", async function (req, res, next){
  const song = await nowPlaying();

  let progress_ms = 0;
  let duration_ms = 0;
  let song_name = "Nothing playing..."
  let song_artist = "";
  let cover_image = "";

  if(song && song["is_playing"]){

    progress_ms = song["progress_ms"];
    duration_ms = song["item"]["duration_ms"];

    song_name = song["item"]["name"]
    song_artist = (song["item"]["artists"] || []).map(({ name }) => name).join(", ");

    const { images = [] } = song["item"]["album"] || {};
    cover_image = images[0]?.url;

  }

  try {
    //Create the Canvas
    const canvas = Canvas.createCanvas(630, 130,"svg");
    const ctx = canvas.getContext("2d");
    ctx.font = "16px RobotoRegular"

    ctx.fillStyle = "transparent"

    let width = canvas.width;

    //Set Song
    const song_text = [
      { text: song_name, fillStyle: "#4e4e4e", font: "21px RobotoBold" }
    ];

    fillMixedTextLeft(ctx, song_text, 130, 52, 500);

    //Set Artist
    const artist_text = [
      { text: song_artist, fillStyle: "#b3b3b3", font: "16px RobotoRegular" }
    ];

    fillMixedTextLeft(ctx, artist_text, 130, 76, 500);

    //Set Progressbar
    if (duration_ms !== 0) {
      ctx.fillStyle = "transparent"
      ctx.strokeStyle = "#666";
      roundRect(ctx, 130, 112, width-140, 8, 2, true);


      let val = Math.ceil((progress_ms / duration_ms * 100 * (width-140)) / 100);

      if (val !== 0) {
        ctx.fillStyle = "#666";
        roundRect(ctx, 130, 112, val, 8, 2, true);
      }
    }




    //Set the Cover
    ctx.strokeStyle = "#ffffff";
    roundRect(ctx,10, 10, 110, 110, 10);

    if(cover_image !== ""){
      const cover = await Canvas.loadImage(cover_image);
      ctx.clip();
      ctx.drawImage(cover, 10, 10, 110, 110);
      ctx.restore();
    }



    res.set("Content-Type", "image/svg+xml");
    res.set("Cache-Control", "s-maxage=1, stale-while-revalidate");
    res.send(canvas.toBuffer());
  }catch (error){
    next(error);
  }
})


//404 Error
app.use((req, res, next) => {
  const err = new Error(`The requested URL was not found on the server.`);
  err.status = 404;
  next(err);
});

//Response the errors
app.use((err, req, res,next) => {
  res.status(err.status || 500);
  res.json({
    error: {
      message: err.message,
    },
    status_code: err.status || 500,
  });
});

app.listen(PORT, () => {
  console.log(
    `Server started on Port ${app.get("port")} | Environment : ${app.get(
      "env"
    )}`
  );
});
