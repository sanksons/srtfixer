### What is upsirshak?

It is a utility program that works on subtitle files [currently only .srt].

### What it does?

Well!! while watching our favourite video series, many times we tend to frustrate seeing the subtitle lagging behind the actual scenes or sometimes they seem to be ahead by a huge margin. This behaviour of subtitle pressures us to do either of the two things:

- Turn Off the subtitle.
- Find a different mathching one on the internet.

upsirshak provides us with a Third option i.e to fix the time of the subtitle file according to your needs.

### How it works?

You need to supply two things to this utility and it will do the rest for you.

- File [path of the .srt file to be fixed]
- Lag or Ahead duration [by what duration subtitle is lagging behind or is Ahead]

### An example is always awesome

```bash
upsirshak -file=/home/Documents/test.srt -sec=1 -msec=1
```



