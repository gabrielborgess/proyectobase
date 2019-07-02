package slices

func RemoveSlice(a []string, i int) []string { // Util para borrar un indice exacto en un slice
	a = append(a[:i], a[i+1:]...)
	return a
}

// Util para agregar elementos extras a un slice con un separador
// Tambien sirve por ejemplo cuando tienes un slice, no quieres hacer otro con strings.split y solo quieres agregar cosas
func AppendSliceBySplit(slice []string, element string, separator string) []string {
	element_slice := strings.Split(element, separator)
	for i := 0; i < len(element_slice); i++ {
		slice = append(slice, element_slice[i])
	}
	return slice
}
