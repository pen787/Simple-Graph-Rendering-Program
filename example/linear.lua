-- Any code outside this will be run when script is load
print("Linear Script load!")
-- Run when graph load
function load()
  print("Loaded!")
  return {R=255,G=0,B=0} -- return color of a graph
end

-- Run for each x coordination
function render(x)
  return x
end