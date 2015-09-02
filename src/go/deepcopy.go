import "reflect"

// Deepcopy recursively deep copies an interface{}
func Deepcopy(iface interface{}) interface{} {
    if iface == nil {
        return nil
    }
    // Make the interface a reflect.Value
    original := reflect.ValueOf(iface)
    // Make a copy of the same type as the original.
    copy := reflect.New(original.Type()).Elem()
    // Recursively copy the original.
    copyRecursive(original, copy)
    // Return theb copy as an interface.
    return copy.Interface()
}

// copyRecursive does the actual copying of the interface. It currently has
// limited support for what it can handle. Add as needed.
func copyRecursive(original, copy reflect.Value) {
    // handle according to original's Kind
    switch original.Kind() {
    case reflect.Ptr:
        // Get the actual value being pointed to.
        originalValue := original.Elem()
        // if  it isn't valid, return.
        if !originalValue.IsValid() {
            return
        }
        copy.Set(reflect.New(originalValue.Type()))
        copyRecursive(originalValue, copy.Elem())
    case reflect.Interface:
        // Get the value for the interface, not the pointer.
        originalValue := original.Elem()
        // Get the value by calling Elem().
        copyValue := reflect.New(originalValue.Type()).Elem()
        copyRecursive(originalValue, copyValue)
        copy.Set(copyValue)
    case reflect.Struct:
        // Go through each field of the struct and copy it.
        for i := 0; i < original.NumField(); i++ {
            copyRecursive(original.Field(i), copy.Field(i))
        }
    case reflect.Slice:
        // Make a new slice and copy each element.
        copy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
        for i := 0; i < original.Len(); i++ {
            copyRecursive(original.Index(i), copy.Index(i))
        }
    case reflect.Map:
        copy.Set(reflect.MakeMap(original.Type()))
        for _, key := range original.MapKeys() {
            originalValue := original.MapIndex(key)
            copyValue := reflect.New(originalValue.Type()).Elem()
            copyRecursive(originalValue, copyValue)
            copy.SetMapIndex(key, copyValue)
        }
    // Set the actual values from here on.
    case reflect.String:
        copy.SetString(original.Interface().(string))

    case reflect.Int:
        copy.SetInt(int64(original.Interface().(int)))

    case reflect.Bool:
        copy.SetBool(original.Interface().(bool))

    case reflect.Float64:
        copy.SetFloat(original.Interface().(float64))

    default:
        copy.Set(original)
    }
}
