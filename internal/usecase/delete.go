package usecase

func (u *Usecase) Delete(key string) error {
	err := u.repository.Delete(key)
	if err != nil {
		return err
	}

	return nil
}
