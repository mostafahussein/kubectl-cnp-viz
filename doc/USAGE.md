## Usage

### Visualize a given network policy

```
kubectl cnp-viz -n my-namespace service-np
```

### Resize the diagram

```
kubectl cnp-viz -n my-namespace service-np --scale 2
```

### Move the diagram horizontally

```
kubectl cnp-viz -n my-namespace service-np --x 600
```

### Move the diagram vertically

```
kubectl cnp-viz -n my-namespace service-np --y 50
```

### Set custom output path

```
kubectl cnp-viz -n my-namespace service-np --output ~/diagrams
```
