package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var r Resource

	for r, err = opener(); err != nil; r, err = opener() {
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}

	defer r.Close()
	defer func() {
		if res := recover(); res != nil {
			switch e := res.(type) {
			case FrobError:
				r.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			}
		}
	}()

	r.Frob(input)
	return err

}
