package erratum

func Use(F ResourceOpener, input string) error {
	resource, err := F()
	defer resource.Close()
	if err != nil {
		return err
	}

	resource.Frob(input)
	resource.Defrob(input)
	return nil
}
