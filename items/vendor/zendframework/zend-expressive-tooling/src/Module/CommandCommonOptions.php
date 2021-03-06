<?php
/**
 * @see       https://github.com/zendframework/zend-expressive-tooling for the canonical source repository
 * @copyright Copyright (c) 2017 Zend Technologies USA Inc. (https://www.zend.com)
 * @license   https://github.com/zendframework/zend-expressive-tooling/blob/master/LICENSE.md New BSD License
 */

declare(strict_types=1);

namespace Zend\Expressive\Tooling\Module;

use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputArgument;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;

/**
 * @internal
 */
final class CommandCommonOptions
{
    /**
     * Add default arguments and options used by all commands.
     *
     * @param Command $command
     */
    public static function addDefaultOptionsAndArguments(Command $command) : void
    {
        $command->addArgument(
            'module',
            InputArgument::REQUIRED,
            $command::HELP_ARG_MODULE
        );

        $command->addOption(
            'composer',
            'c',
            InputOption::VALUE_REQUIRED,
            'Specify the path to the composer binary; defaults to "composer"'
        );

        $command->addOption(
            'modules-path',
            'p',
            InputOption::VALUE_REQUIRED,
            'Specify the path to the modules directory; defaults to "src"'
        );
    }

    /**
     * Retrieve the modules path from input
     */
    public static function getModulesPath(InputInterface $input) : string
    {
        $modulesPath = $input->getOption('modules-path') ?: 'src';
        $modulesPath = preg_replace('/^\.\//', '', str_replace('\\', '/', $modulesPath));
        return $modulesPath;
    }
}
