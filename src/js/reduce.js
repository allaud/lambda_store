function reduce(array, base, combine) {
  var forEach = function (array, action) {
    for (var i = 0; i < array.length; i++)
      action(array[i]);
  }

  forEach(array, function (element) {
    base = combine(base, element);
  });
  return base;
}

