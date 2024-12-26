BEGIN;

CREATE TABLE IF NOT EXISTS public.ledgers
(
    id         BIGINT PRIMARY KEY CHECK (id > 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE NOT NULL,
    name       VARCHAR(35)              NOT NULL CHECK (name <> ''),
    currency   VARCHAR(3)               NOT NULL CHECK (currency <> '')
);
COMMENT ON COLUMN public.ledgers.id IS 'Snowflake ID';
COMMENT ON COLUMN public.ledgers.currency IS 'Currency in ISO format';

CREATE TABLE IF NOT EXISTS public.access
(
    id          BIGINT PRIMARY KEY CHECK (id > 0),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at  TIMESTAMP WITH TIME ZONE,
    ledger_id   BIGINT REFERENCES public.ledgers (id),
    user_id     VARCHAR(50)              NOT NULL CHECK (user_id <> '') ,
    access_type VARCHAR(15)              NOT NULL CHECK (access_type <> '')
);
CREATE INDEX IF NOT EXISTS access_user_id_idx ON public.access (user_id);
COMMENT ON COLUMN public.access.id IS 'Snowflake ID';
COMMENT ON COLUMN public.access.user_id IS 'User ID from the IAM system';

CREATE TABLE IF NOT EXISTS public.accounts
(
    id           BIGINT PRIMARY KEY CHECK (id > 0),
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    ledger_id    BIGINT REFERENCES public.ledgers (id),
    name         VARCHAR(50)              NOT NULL CHECK (name <> ''),
    account_type VARCHAR(15)              NOT NULL CHECK (account_type <> '')
);
COMMENT ON COLUMN public.accounts.id IS 'Snowflake ID';

CREATE TABLE IF NOT EXISTS public.transaction_item_categories
(
    id         BIGINT PRIMARY KEY CHECK (id > 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE NOT NULL,
    ledger_id  BIGINT REFERENCES public.ledgers (id),
    name       VARCHAR(50)              NOT NULL CHECK (name <> '')
);
COMMENT ON COLUMN public.transaction_item_categories.id IS 'Snowflake ID';

CREATE TABLE IF NOT EXISTS public.transaction_items
(
    id          BIGINT PRIMARY KEY CHECK (id > 0),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at  TIMESTAMP WITH TIME ZONE,
    ledger_id   BIGINT REFERENCES public.ledgers (id),
    category_id BIGINT REFERENCES public.transaction_item_categories (id),
    name        VARCHAR(50)              NOT NULL CHECK (name <> ''),
    item_type   VARCHAR(15)              NOT NULL CHECK (item_type <> '')
);
COMMENT ON COLUMN public.transaction_items.id IS 'Snowflake ID';

CREATE TABLE IF NOT EXISTS public.payees
(
    id         BIGINT PRIMARY KEY CHECK (id > 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    ledger_id  BIGINT REFERENCES public.ledgers (id),
    name       VARCHAR(50)              NOT NULL CHECK (name <> ''),
    payee_type VARCHAR(15)              NOT NULL CHECK (payee_type <> '')
);
COMMENT ON COLUMN public.payees.id IS 'Snowflake ID';

CREATE TABLE IF NOT EXISTS public.transactions
(
    id               BIGINT PRIMARY KEY CHECK (id > 0),
    created_at       TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at       TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at       TIMESTAMP WITH TIME ZONE,
    ledger_id        BIGINT REFERENCES public.ledgers (id),
    item_id          BIGINT REFERENCES public.transaction_items (id),
    txn_date         TIMESTAMP WITH TIME ZONE NOT NULL,
    source_type      VARCHAR(15)              NOT NULL CHECK (source_type <> ''),
    source_id        BIGINT                   NOT NULL CHECK (source_id > 0),
    destination_type VARCHAR(15)              NOT NULL CHECK (destination_type <> ''),
    destination_id   BIGINT                   NOT NULL CHECK (source_id > 0),
    amount           INT                      NOT NULL CHECK (amount <> 0)
);
COMMENT ON COLUMN public.transactions.id IS 'Snowflake ID';
COMMENT ON COLUMN public.transactions.source_type IS 'Either Payee or Account';
COMMENT ON COLUMN public.transactions.destination_type IS 'Either Payee or Account';

CREATE TABLE IF NOT EXISTS public.transaction_item_budgets
(
    id          BIGINT PRIMARY KEY CHECK (id > 0),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at  TIMESTAMP WITH TIME ZONE,
    ledger_id   BIGINT REFERENCES public.ledgers (id),
    item_id     BIGINT REFERENCES public.transaction_items (id),
    budget_date TIMESTAMP WITH TIME ZONE NOT NULL,
    budgeted    INT                      NOT NULL CHECK (budgeted >= 0),
    transacted  INT                      NOT NULL
);
COMMENT ON COLUMN public.transaction_item_budgets.id IS 'Snowflake ID';
COMMENT ON COLUMN public.transaction_item_budgets.budget_date IS 'Month of the year';

CREATE TABLE IF NOT EXISTS public.budget_plans
(
    id         BIGINT PRIMARY KEY CHECK (id > 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    ledger_id  BIGINT REFERENCES public.ledgers (id),
    name       VARCHAR(50)              NOT NULL CHECK (name <> '')
);
COMMENT ON COLUMN public.budget_plans.id IS 'Snowflake ID';

CREATE TABLE IF NOT EXISTS public.budget_plan_items
(
    id         BIGINT PRIMARY KEY CHECK (id > 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    ledger_id  BIGINT REFERENCES public.ledgers (id),
    plan_id    BIGINT REFERENCES public.budget_plans (id),
    item_id    BIGINT REFERENCES public.transaction_items (id),
    amount     INT                      NOT NULL CHECK (amount >= 0)
);
COMMENT ON COLUMN public.budget_plan_items.id IS 'Snowflake ID';

COMMIT;
