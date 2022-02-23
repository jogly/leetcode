package main

func isNumber(s string) bool {
	n, ok := number([]byte(s))
	return ok && n == len(s)
}

type consumer func([]byte) (int, bool)

var (
	e    = symbol([]byte{'e', 'E'})
	sign = symbol([]byte{'-', '+'})
	dot  = symbol([]byte{'.'})

	oneOrMoreDigits = min(consume(digit), 1)

	parseInt = sequence(
		optional(sign),
		oneOrMoreDigits,
	)

	parseDec = sequence(
		optional(sign),
		oneOf(
			sequence(oneOrMoreDigits, dot, oneOrMoreDigits),
			sequence(oneOrMoreDigits, dot),
			sequence(dot, oneOrMoreDigits),
		),
	)
	trailer = sequence(optional(sequence(e, parseInt)), end)
	number  = oneOf(
		sequence(parseDec, trailer),
		sequence(parseInt, trailer),
	)
)

func end(b []byte) (int, bool) {
	return 0, len(b) == 0
}

func optional(c consumer) consumer {
	return func(b []byte) (int, bool) {
		if n, ok := c(b); ok {
			return n, true
		}
		return 0, true
	}
}

func oneOf(cs ...consumer) consumer {
	return func(b []byte) (int, bool) {
		for _, c := range cs {
			if n, ok := c(b); ok {
				return n, ok
			}
		}
		return 0, false
	}
}

func digit(b []byte) (int, bool) {
	if len(b) == 0 {
		return 0, false
	}
	if b[0] >= '0' && b[0] <= '9' {
		return 1, true
	}
	return 0, false
}

func symbol(syms []byte) consumer {
	return func(b []byte) (int, bool) {
		if len(b) == 0 {
			return 0, false
		}
		for _, s := range syms {
			if s == b[0] {
				return 1, true
			}
		}
		return 0, false
	}
}

func min(f consumer, min int) consumer {
	return func(b []byte) (int, bool) {
		n, ok := f(b)
		return n, ok && n >= min
	}
}

func sequence(cs ...consumer) consumer {
	return func(b []byte) (int, bool) {
		var total, n int
		var ok bool
		for _, c := range cs {
			n, ok = c(b[total:])
			if !ok {
				return total, ok
			}
			total += n
		}
		return total, ok
	}
}

func consume(c consumer) consumer {
	return func(b []byte) (int, bool) {
		n, ok := c(b)
		total := n
		for ok && total < len(b) {
			n, ok = c(b[total:])
			if n == 0 && ok {
				panic("you piece of shit dont do that")
			}
			total += n
		}
		return total, true
	}
}
