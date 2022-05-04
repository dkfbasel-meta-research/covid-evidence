package ictrp

// UpdateScreening will update the screening variable in the screening table
func (who *ICTRP) UpdateScreening() error {
	return who.Store.SetScreening("ictrp")
}
