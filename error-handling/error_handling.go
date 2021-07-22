package erratum

import "log"

func Use(opener ResourceOpener, input string) error {
	r, err := opener()
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred: ", err)
		} else {
			r.Close()
		}
	}()

	switch err.(type) {
	case *TransientError:
		Use(opener, input)
	case nil:
		break
	default:
		return err
	}

	r.Frob(input)
	r.Defrob(input)
	return nil
}
