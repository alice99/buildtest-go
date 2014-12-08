package x

func X() string { return "x" }

func Y() string { return X() }

var Z2 = X()
