# go-enumerable

This is enumerable library inspired by [.NET Enumerable](https://docs.microsoft.com/en-us/dotnet/api/system.linq.enumerable?view=netcore-3.1).

This library is using [go2 generics proposal/type parameters](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md), so you can use this only in go2go environment.
To setup go2go environment, please refer https://github.com/golang/go/issues/39619.

# Usage

Currently(2020/08/11) go2go does not support go modules, so you can try in go-enumerable test.

```go
package enumerable

import "testing"

func TestExample(t *testing.T) {
	equal(t, []int{40, 60}, ToSlice(
		Where(
			Select(
				Slice([]int{10, 20, 30}),
				func(v int) int { return v * 2 },
			),
			func(v int) bool { return v > 20 },
		),
	))
}
```

This is almost same below C# code using LINQ.

```cs
using System;
using System.Collections.Generic;
using System.Linq;

public class Program
{
	public static void Main()
	{
		var list = new List<int>(){10, 20, 30};
		list.Select(v => v * 2).Where(v => v > 20).ToList(); // => List<int>(){ 20, 40 }
	}
}
```

# FAQ

Q. Why not method chain?

A. Method type parameter must be declared in struct/interface in draft spec(https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md#methods-may-not-take-additional-type-arguments), so we can not define mapping method like `Select`.

