INSERT OR IGNORE INTO ingredients(name) VALUES
('egg'),('spinach'),('olive oil'),('salt'),('pepper'),
('chicken'),('lettuce'),('tomato'),('cucumber'),('lemon'),
('pasta'),('garlic'),('basil'),('parmesan'),('oats'),
('milk'),('honey'),('chia'),('lentil'),('onion'),
('carrot'),('celery'),('cumin');

INSERT INTO recipes(title, description, meal_type, cook_time_minutes) VALUES
('Spinach Omelette','Fluffy omelette with fresh spinach','breakfast',10),
('Chicken Salad','Light salad with lemon dressing','lunch',15),
('Pasta Pomodoro','Simple tomato-basil pasta','dinner',20),
('Overnight Oats','No-cook creamy oats','breakfast',5),
('Lentil Soup','Hearty, spiced lentil soup','dinner',35);

-- Link ingredients via names
INSERT OR IGNORE INTO recipe_ingredients(recipe_id, ingredient_id, quantity)
SELECT r.id, i.id, '2' FROM recipes r JOIN ingredients i ON i.name='egg' WHERE r.title='Spinach Omelette';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1 cup' FROM recipes r JOIN ingredients i ON i.name='spinach' WHERE r.title='Spinach Omelette';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1 tbsp' FROM recipes r JOIN ingredients i ON i.name='olive oil' WHERE r.title='Spinach Omelette';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, 'pinch' FROM recipes r JOIN ingredients i ON i.name='salt' WHERE r.title='Spinach Omelette';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, 'pinch' FROM recipes r JOIN ingredients i ON i.name='pepper' WHERE r.title='Spinach Omelette';

INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '150 g' FROM recipes r JOIN ingredients i ON i.name='chicken' WHERE r.title='Chicken Salad';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1 cup' FROM recipes r JOIN ingredients i ON i.name='lettuce' WHERE r.title='Chicken Salad';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1' FROM recipes r JOIN ingredients i ON i.name='tomato' WHERE r.title='Chicken Salad';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1/2' FROM recipes r JOIN ingredients i ON i.name='cucumber' WHERE r.title='Chicken Salad';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1 tbsp' FROM recipes r JOIN ingredients i ON i.name='olive oil' WHERE r.title='Chicken Salad';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1 tbsp' FROM recipes r JOIN ingredients i ON i.name='lemon' WHERE r.title='Chicken Salad';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, 'pinch' FROM recipes r JOIN ingredients i ON i.name='salt' WHERE r.title='Chicken Salad';

INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '150 g' FROM recipes r JOIN ingredients i ON i.name='pasta' WHERE r.title='Pasta Pomodoro';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '2 cloves' FROM recipes r JOIN ingredients i ON i.name='garlic' WHERE r.title='Pasta Pomodoro';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '5' FROM recipes r JOIN ingredients i ON i.name='tomato' WHERE r.title='Pasta Pomodoro';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, 'handful' FROM recipes r JOIN ingredients i ON i.name='basil' WHERE r.title='Pasta Pomodoro';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, 'to serve' FROM recipes r JOIN ingredients i ON i.name='parmesan' WHERE r.title='Pasta Pomodoro';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1 tbsp' FROM recipes r JOIN ingredients i ON i.name='olive oil' WHERE r.title='Pasta Pomodoro';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, 'pinch' FROM recipes r JOIN ingredients i ON i.name='salt' WHERE r.title='Pasta Pomodoro';

INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1/2 cup' FROM recipes r JOIN ingredients i ON i.name='oats' WHERE r.title='Overnight Oats';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1/2 cup' FROM recipes r JOIN ingredients i ON i.name='milk' WHERE r.title='Overnight Oats';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1 tsp' FROM recipes r JOIN ingredients i ON i.name='honey' WHERE r.title='Overnight Oats';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1 tsp' FROM recipes r JOIN ingredients i ON i.name='chia' WHERE r.title='Overnight Oats';

INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1 cup' FROM recipes r JOIN ingredients i ON i.name='lentil' WHERE r.title='Lentil Soup';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1' FROM recipes r JOIN ingredients i ON i.name='onion' WHERE r.title='Lentil Soup';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1' FROM recipes r JOIN ingredients i ON i.name='carrot' WHERE r.title='Lentil Soup';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1' FROM recipes r JOIN ingredients i ON i.name='celery' WHERE r.title='Lentil Soup';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, '1 tsp' FROM recipes r JOIN ingredients i ON i.name='cumin' WHERE r.title='Lentil Soup';
INSERT OR IGNORE INTO recipe_ingredients SELECT r.id, i.id, 'pinch' FROM recipes r JOIN ingredients i ON i.name='salt' WHERE r.title='Lentil Soup';
