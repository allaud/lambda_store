function pick(obj, list, context) {
  var result = {};

  if (typeof list === 'string') {
    list = [list];
  }

  Object.keys(obj)
    .forEach(function(key) {
      if (list.indexOf(key) > -1) {
        result[key] = obj[key];
      }
    }, context);

  return result;
}
