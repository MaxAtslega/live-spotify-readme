/**Create align right mixed text
 * @param {number} x - x coordinate
 * @param {number} y - y coordinate
 * @param {number} maxWidth - max width
 * @param {Object[]} args - List of Text
 * @param {any} ctx - Canvas Context
 * */
module.exports = (ctx, args, x, y, maxWidth) => {
  let defaultFillStyle = ctx.fillStyle;
  let defaultFont = ctx.font;
  ctx.textAlign = "start";
  ctx.save();


  args.forEach(({ text, fillStyle, font }) => {
    ctx.fillStyle = fillStyle || defaultFillStyle;
    ctx.font = font || defaultFont;
    ctx.fillText(fittingString(ctx,text, maxWidth), x, y);
    x += ctx.measureText(text).width;
  });
  ctx.restore();
};

function fittingString(c, str, maxWidth) {
  var width = c.measureText(str).width;
  var ellipsis = '…';
  var ellipsisWidth = c.measureText(ellipsis).width;
  if (width<=maxWidth || width<=ellipsisWidth) {
    return str;
  } else {
    var len = str.length;
    while (width>=maxWidth-ellipsisWidth && len-->0) {
      str = str.substring(0, len);
      width = c.measureText(str).width;
    }
    return str+ellipsis;
  }
}