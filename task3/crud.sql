-- 题目1：
-- 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
-- 1. 插入
INSERT INTO students (name, age, grade) VALUES ('张三', 20, '三年级');
-- 2. 查询
SELECT * FROM students WHERE age > 18;
-- 3. 更新
UPDATE students SET grade = '四年级' WHERE name = '张三';
-- 4. 删除
DELETE FROM students WHERE age < 15;