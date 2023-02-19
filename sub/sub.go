package sub

// Sub this is sub func.
func Sub(stopAllCh chan struct{}) {
	go subA(stopAllCh)
	go subB(stopAllCh)
	go subC(stopAllCh)
	go subD(stopAllCh)
}
