-- 题目2：
-- 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
-- 使用事务，假设在应用程序中控制条件：

START TRANSACTION;

-- 锁定账户A的行
SELECT balance FROM accounts WHERE id = 1 FOR UPDATE;

-- 应用程序检查余额，如果余额>=100，则执行：
UPDATE accounts SET balance = balance - 100 WHERE id = 1;
UPDATE accounts SET balance = balance + 100 WHERE id = 2;
INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (1, 2, 100);
COMMIT;

-- 否则：
ROLLBACK;