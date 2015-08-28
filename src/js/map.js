function map(array, func) {
  var forEach = function (array, action) {
    for (var i = 0; i < array.length; i++)
      action(array[i]);
  }
  var result = [];
  forEach(array, function (element) {
    result.push(func(element));
  });
  return result;
}
