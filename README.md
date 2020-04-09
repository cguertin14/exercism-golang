# Exercism Go Track

My solutions to the coding exercices from exercism.

## Syntax Conventions

In Go this is not a matter of style, it is a matter of syntax.

Exported names (that is, identifiers that can be used from a package other than the one where they are defined) begin with a capital letter. Thus if your method is part of your public API, it should be written:

`WriteToDB`

but if it is an internal helper method it should be written:

`writeToDB`

The benefit of doing it this way over using keywords to define exportedness (`extern`, `public`, etc.) is that making it a part of the name ensures that anywhere an identifier is used you can tell if it is exported or not without having to find where it was defined (to see if the definition contains a keyword).

**See also:** [*Exported Identifiers*](https://golang.org/ref/spec#Exported_identifiers) from the spec.

### i18n

Because Go is UTF-8 encoded and supports any Unicode character with the letters or numers property in identifier names some people in locales that don't have a concept of case may have trouble creating exported methods (the default is non-exported). In this case (pun intended) it is common to prefix identifiers with an `X` to indicate exportedness. For example: `X日本語`

**See also:** [*What's up with Unicode identifiers?*](https://golang.org/doc/faq#unicode_identifiers) from the FAQ.

## Style

As far as the general style goes, it is to always use camel-case (except for the first letter, as previously mentioned). This includes constants, functions, and other identifiers. So for example a list of (exported) constants might look like:

<!-- language: go-lang -->

    const (
        StateConnected = iota
        StateError
        StateDone

        internalStateMask = 0x2 
    )

Furthermore, abbreviations are always written with the same case, so you would write one of the following:

    dbWrite
    writeDB

instead of `writeDb` or `DbWrite`.

## References

[Syntax](https://stackoverflow.com/a/38617771/9201984)
