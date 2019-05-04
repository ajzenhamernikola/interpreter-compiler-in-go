# Monkey interpretator i kompilator

Ovaj repozitorijum sadrži izvorni kod za interpretator i kompilator 
za programski jezik Monkey. Oba proizvoda su napisana u programskom
jeziku Go.

Za više detalja o proizvodima, pogledati README stranice posvećene 
njima:

- [Interpretator](interpreter/README.md)
- Kompilator

## Potrebne instalacije

1. [Go 1.12.4](https://golang.org/dl/)
    - [Brzi download link za Windows OS (118MB)](https://dl.google.com/go/go1.12.4.windows-amd64.msi)
    - [Brzi download link za macOS (121MB)](https://dl.google.com/go/go1.12.4.darwin-amd64.pkg)
    - [Brzi download link za Linux (122MB)](https://dl.google.com/go/go1.12.4.linux-amd64.tar.gz)
 
## Programski jezik Monkey

Primer izvornog koda napisanog u programskom jeziku Monkey 
koji ilustruje njegove mogućnosti:

```
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);

let myArray = [1, 2, 3, 4, 5];

let thorsten = {"name": "Thorsten", "age": 28};

myArray[0] // => 1
thorsten["name"] // => "Thorsten"

let add = fn(a, b) { 
    return a + b; 
};

add(1, 2); // => 3

let add2 = fn(a, b) { 
    a + b; 
};

add2(1, 2); // => 3

let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};

let twice = fn(f, x) {
    return f(f(x));
};

let addTwo = fn(x) {
    return x + 2;
};

twice(addTwo, 2); // => 6
```