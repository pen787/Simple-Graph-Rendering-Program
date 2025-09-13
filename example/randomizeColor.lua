-- Run when graph load
function load()
  local red = math.random(0, 255)
  local green = math.random(0, 255)
  local blue = math.random(0, 255)
  
  return {R=red,G=green,B=blue} -- return color of a graph
end

-- Run for each x coordination
function render(x)
  return 1/x
end