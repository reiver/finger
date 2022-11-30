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

So we see that, `finger` can be called with:

* an optional _switch_,
* an optional _user_, and
* zero, one, or many _@host_

Here are a bunch of examples to try to help show what this really means:
```
finger
finger /W
finger /W joeblow
finger /W joeblow@example.com
finger /W joeblow@example.com@changelog.ca
finger /W @example.com
finger /W @example.com@changelog.ca
finger joeblow
finger joeblow@example.com
finger joeblow@example.com@changelog.ca
finger @example.com
finger @example.com@changelog.ca
finger dariush@once.com@twice.net@thrice.org@fource.net
finger /PULL charlie@changelog.ca
finger /PUSH malekeh@changelog.ca
finger /LIST dariush@changelog.ca
finger /PICK elizabeth@changelog.ca
finger /path/to/something.ext
finger /path/to/something.ext joeblow@example.com
finger /path/to/something.ext joeblow@example.com@changelog.ca
```
