function extend() {
  var args = Array.prototype.slice.call(arguments);
  var orig = args.shift();
  if (args.length === 0) {
    return orig;
  }
  for(var i=0;i<=args.length;i++){
    var r = args[i];
    for (var key in r) {
      if(Object.prototype.hasOwnProperty.call(r, key)) {
        orig[key] = r[key];
      }
    }
  }
  return orig;
}
