# Monkey interpretator

Ovaj direktorijum sadrži izvorni kod za implementiranje interpretatora
za programski jezik Monkey. Izvorni kod je podeljen na pakete na
sledeći način:

- Leksička analiza:
    - [rad sa tokenima](src/token)
    - [implementacija leksera](src/lexer)

Dodatni paketi koji su uključeni su:

- Interpreter:
    - [implementacija REPL-a](src/repl)
    - [ulazni/glavni deo implementacije](src/main)