def benchmark(func):
    "wrapper for function benchmarking"
    from time import time
    def wrapper(*args, **kwargs):
        start = time()
        result = func(*args, **kwargs)
        difference = (time() - start) * 1000  # milliseconds
        message = 'Benchmark: #' + func.__name__
        message += ': ' + ("%0.2f" % difference)
        return result
    return wrapper
