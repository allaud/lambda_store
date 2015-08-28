function partial(func) {
  var asArray = function(quasiArray, start) {
    var result = [];
    for (var i = (start || 0); i < quasiArray.length; i++)
      result.push(quasiArray[i]);
    return result;
  }

  var fixedArgs = asArray(arguments, 1);
  return function(){
    return func.apply(null, fixedArgs.concat(asArray(arguments)));
  };
}
