def lam(n):
    """ this is just a lambda func'
     just a test'
    """
    return lambda m: m + n
    
t = lam(5)

print(t(6))
print(lam.__doc__)