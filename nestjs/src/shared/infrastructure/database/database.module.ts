import { DynamicModule, Global, Module } from "@nestjs/common";
import { PrismaClient } from "@prisma/client";
import { ConfigService } from "@nestjs/config";
import { EnvConfigModule } from "../env-config/env-config.module";
import { PrismaService } from "@/booking/adapters/prisma.service";

@Global()
@Module({
    imports: [EnvConfigModule.forRoot()],
    providers: [PrismaService, ConfigService],
    exports: [PrismaService]
})
export class DatabaseModule {
    static forTest(prismaClient: PrismaClient): DynamicModule {
        return {
            module: DatabaseModule,
            providers: [
                {
                    provide: PrismaService,
                    useFactory: ()=> prismaClient as PrismaService,
                }
            ]
        }
    }
}