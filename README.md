# Monkey interpreter i kompilator

Ovaj repozitorijum sadrži izvorni kod za interpreter i kompilator 
za programski jezik Monkey. Oba proizvoda su napisana u programskom
jeziku Go.

Za više detalja o proizvodima, pogledati README stranice posvećene 
njima:

- [Interpreter](interpreter/)
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
// Dinamicki tipizirane promenljive, tipovi podataka:

// Celi brojevi
let age = 1; 
// Niske
let name = "Monkey"; 
// Aritmeticki izrazi
let result = 10 * (20 / 2); 
// Nizovi
let myArray = [1, 2, 3, 4, 5];
// Recnici
let thorsten = {"name": "Thorsten", "age": 28};

// Pristupanje elementima nizova i recnika
myArray[0] // => 1
thorsten["name"] // => "Thorsten"

// Definisanje funkcija
let add = fn(a, b) { 
    return a + b; 
};

// Pozivanje funkcija
add(1, 2); // => 3

let add2 = fn(a, b) {
    // Podrazumevane povratne vrednosti 
    a + b; 
};

add2(1, 2); // => 3

// Rekurzivne funkcije
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

// Funkcije viseg reda
let twice = fn(f, x) {
    return f(f(x));
};

let addTwo = fn(x) {
    return x + 2;
};

twice(addTwo, 2); // => 6
```