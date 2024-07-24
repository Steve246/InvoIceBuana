-- Drop foreign keys only if they exist
-- ALTER TABLE InvoiceItem DROP FOREIGN KEY IF EXISTS invoiceitem_ibfk_1;
-- ALTER TABLE InvoiceItem DROP FOREIGN KEY IF EXISTS invoiceitem_ibfk_2;
-- ALTER TABLE Invoice DROP FOREIGN KEY IF EXISTS invoice_ibfk_1;

-- Drop tables only if they exist
DROP TABLE IF EXISTS InvoiceItem;
DROP TABLE IF EXISTS Invoice;
DROP TABLE IF EXISTS Customer;
DROP TABLE IF EXISTS Item;
