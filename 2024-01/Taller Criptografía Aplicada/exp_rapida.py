import numpy as np

def exp_rap_binario(a,b,n):
    x = 1
    for i in range(len(b)):
        if b[i] == "1":
            x = (np.power(x,2) * a) % n
        else:
            x = np.power(x,2) % n
    return x

a = 13
b_binario = "100000001"
n = 27

print("El resultado obtenido es: ",exp_rap_binario(a,b_binario,n))



