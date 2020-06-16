#!/usr/bin/env python
from pathlib import Path

if __name__ == '__main__':
    cur_path = Path(__file__).parent
    utils_path = cur_path.joinpath('..', 'internal', 'utils')
    utils_path = utils_path.resolve()
    with utils_path.joinpath('filter.go').open('w') as fp:
        fp.write('package utils\n')
        for t in ['int', 'float64', 'string']:
            tt = t.title()
            fp.write(
                f"""
// Filter{tt} function
func Filter{tt}(arr []{t}, cb func({t}) bool) []{t} {{
	rv := arr[:0]
	for _, v := range arr {{
		if cb(v) {{
			rv = append(rv, v)
		}}
	}}
	return rv
}}
"""
            )
