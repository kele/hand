# Could You Give Your Program a Helping Hand From Time to Time?

Imagine you process lots of data. From time to time, though, the data entries
are a bit tricky to handle. Writing code that handles them would require a lot
of effort... but you could **easily** handle these individual cases yourself,
because you are a human.

The only thing you need to do to enable this is to change your code from looking
like so:

```
if v, err := Convert("shirt mint"); err != nil {
    return nil, err
} 
process(v)
```

to

```
if v, err := hand.HelpWith(Convert)("shirt mint"); err != nil {
    return nil, err
} 
process(v)
```

If the `Convert()` function were to return an error, a prompt in the terminal
like so appears:

```
hand: /home/kele/hand/hand.go:20 -- f([shirt mint]) = (_, unknown color category).
hand: Fix?
```

and you are expected to supply a JSON value and send `EndOfFile` (Ctrl+D in most
terminals). That value will be now returned instead of the error.