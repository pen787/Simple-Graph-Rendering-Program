-- Any code outside this will be run when script is load
print("log Script load!")
-- Run when graph load
function load()
  print("Loaded!")
  return {R=255,G=100,B=50} -- return color of a graph
end

-- Run for each x coordination
function render(x)
  return math.log(x, 10)
end