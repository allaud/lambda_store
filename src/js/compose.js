function compose(func1, func2) {
  return function() {
    return func1(func2.apply(null, arguments));
  };
}
