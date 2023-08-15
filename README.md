# Out2JSON
Simple CLI utility that converts a text-based outline format to JSON.

For now, it assumes indentation level is four spaces.

For example, given the following input:

```
Root
- Child 1
    - Grandchild 1
    - Grandchild 2
- Child 2
```

It would output:
```json
{
    "text": "Root",
    "depth": 0,
    "children": [
      {
        "text": "Child 1",
        "depth": 1,
        "children": [
          {
            "text": "Grandchild 1",
            "depth": 2,
            "children": []
          },
          {
            "text": "Grandchild 2",
            "depth": 2,
            "children": []
          }
        ]
      },
      {
        "text": "Child 2",
        "depth": 1,
        "children": []
      }
    ]
}
```
