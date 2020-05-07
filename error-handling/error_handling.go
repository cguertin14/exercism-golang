package erratum

import "fmt"

// Use is a function that accepts a ResourceOpener and a string as parameters.
func Use(o ResourceOpener, input string) (e error) {
	res, err := o()

	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		res, err = o()
	}

	defer res.Close()

	defer func() {
		if recErr := recover(); recErr != nil {
			if frob, ok := recErr.(FrobError); ok {
				res.Defrob(frob.defrobTag)
			}
			e = fmt.Errorf("%v", recErr)
		}
	}()

	res.Frob(input)

	return nil
}
