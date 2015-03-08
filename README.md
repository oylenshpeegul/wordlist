# wordlist
[Wordlist](http://oylenshpeegul.github.io/Compress-DAWG/) compression in Go.

This is a compression scheme for sorted lists of words. I learned it
years ago from my friend Mike, who did it in Fortran. I re-did it
[in Perl](https://github.com/oylenshpeegul/Compress-DAWG) and again
[in Python](https://github.com/oylenshpeegul/Compress-DAWG/blob/master/wordlist.py). Now I've done it in Go.

I have since learned that this scheme is called
[front compression](http://en.wikipedia.org/wiki/Incremental_encoding)
and it's fairly common.
* [Crack](http://www.crypticide.com/alecm/software/crack/c50-faq.html) dictionaries.
* [word-list-compress](http://www.man-online.org/page/1-word-list-compress/)
used by [aspell](http://aspell.net/).
* [Unix locate](http://www.eecs.berkeley.edu/Pubs/TechRpts/1983/CSD-83-148.pdf)(pdf)...thanks, [@b0rk](http://jvns.ca/blog/2015/03/05/how-the-locate-command-works-and-lets-rewrite-it-in-one-minute/)!

