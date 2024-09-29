/*
  Warnings:

  - You are about to drop the column `Status` on the `booking` table. All the data in the column will be lost.
  - Added the required column `documentType` to the `guest` table without a default value. This is not possible if the table is not empty.
  - Added the required column `idNumber` to the `guest` table without a default value. This is not possible if the table is not empty.
  - Added the required column `currency` to the `room` table without a default value. This is not possible if the table is not empty.
  - Added the required column `value` to the `room` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "booking" DROP COLUMN "Status",
ADD COLUMN     "status" INTEGER NOT NULL DEFAULT 0;

-- AlterTable
ALTER TABLE "guest" ADD COLUMN     "documentType" INTEGER NOT NULL,
ADD COLUMN     "idNumber" TEXT NOT NULL;

-- AlterTable
ALTER TABLE "room" ADD COLUMN     "currency" INTEGER NOT NULL,
ADD COLUMN     "value" DOUBLE PRECISION NOT NULL;
