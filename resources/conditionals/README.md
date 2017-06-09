# If

Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

## If with a short statement

Like for, the if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.

(Try using v in the last return statement.)

## If and else

Variables declared inside an if short statement are also available inside any of the else blocks.

## Switch

A case body breaks automatically, unless it ends with a fallthrough statement.
Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

For example,

    switch i {
        case 0:
        case f():
    }

does not call f if i==0.

## Switch with no condition

Switch without a condition is the same as switch true.

This construct can be a clean way to write long if-then-else chains.