# EdMe
>Источник -> [Stackoverflow](https://ru.stackoverflow.com/q/548852)

```go
secret := rand.Intn(10) + 1
```

- `secret` - записываем действия в переменную
- `rand.Intn` - генерирует число от 0 до 9
- `+ 1` - сдвигаем диапазон на 1, теперь с 1 до 10
### Угадал ли пользователь

```go
    var consoleInput int  
    for {  
       fmt.Print("Ваш вариант: ")  
       fmt.Scanln(&consoleInput)  
  
       if consoleInput == secret {  
          fmt.Printf("Вы угадали! Это %d.\n", secret)  
          break  
       } else if consoleInput < secret {  
          fmt.Println("Загаданное число больше!")  
       } else {  
          fmt.Println("Загаданное число меньше!")  
       }  
    }  
}
```

- `var guess int` и `fmt.Scanln(&guess)` - считывание и записывание числа с консоли
- `if guess == secret` - если  число от пользователя (`consoleInput`) оказывается числом которое сгенерила  `secret` через, то выводит что пользователь прав, а дальше методом исключения
# Что можно изменить

- `rand.Intn` на `rand.Float64` - переводит в float64
---
