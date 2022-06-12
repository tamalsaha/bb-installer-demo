# generator

## sample code

```
	out.Version = in.Version
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.ConfigSecret = (*v1.LocalObjectReference)(unsafe.Pointer(in.ConfigSecret))
	out.DataSources = *(*[]v1.VolumeSource)(unsafe.Pointer(&in.DataSources))
	out.TLS = (*apiv1.TLSConfig)(unsafe.Pointer(in.TLS))
	if err := Convert_v1alpha1_BackendStorageSpec_To_v1alpha2_BackendStorageSpec(&in.Backend, &out.Backend, s); err != nil {
		return err
	}
	if in.Unsealer != nil {
		in, out := &in.Unsealer, &out.Unsealer
		*out = new(v1alpha2.UnsealerSpec)
		if err := Convert_v1alpha1_UnsealerSpec_To_v1alpha2_UnsealerSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Unsealer = nil
	}
```
