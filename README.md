# finger

**finger** is a modern **finger-protocol** **client**.

The **fingerverse** (i.e., the collection of all finger-sites) is an early Internet-based social networks, with its origins in the early 1970s.
Like with a lot of things from the early Internet — the fingervese is a decentralized network.

## Try It Out

Once you have `finger` installed, try running this command:
```
finger reiver@plan.cat
```

Although there are other ways you can use `finger`— in most case, you are probably going to use `finger` in this way.
But just with a different query.

For example, here is a similar example usage of `finger` but with a different query (of the same `user@host` type):
```
finger charles@finger.farm
```

## Finger Query

The way you can understand a finger-query is:
```
    finger joeblow@example.com
           \_____/ \_________/
             |           |
         a person      a community
             or           named
         a service     “example.com”
           named
         “joeblow”
```

So with this query —

* there is a community (running on a computer connected to the Internet) called “`example.com`”, and
* there is a person (within that community) whose username is “`joeblow`”.

## Usage

Here is how, in general, you can use `finger`:

```
finger [/switch] [user][@host…]
```
